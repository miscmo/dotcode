#include <vector>
#include <iostream>
#include <queue>
#include <thread>
#include <functional>
#include <mutex>
#include <condition_variable>
#include <assert.h>

#include <unistd.h>

using namespace std;

#define THREADPOOL_MAX_NUM 16

class ThreadPool{
public:
    enum taskPriorityE { level0, level1, level2, };
    typedef std::function<void()> Task;
    typedef std::pair<taskPriorityE, Task> TaskPair;    // 1. 任务优先级可以使用优先队列（堆实现）

    ThreadPool(unsigned int threadSize = 4) 
        : m_mutex() 
        , m_cond()
        , m_isStarted(true) {
        assert(m_threads.empty());
        m_isStarted = true;
        for (int i = 0; i < threadSize; ++i)
        {
            m_threads.push_back(new std::thread(std::bind(&ThreadPool::threadLoop, this)));
        }
    }

    ~ThreadPool() {
        if(m_isStarted) {
            {
                std::unique_lock<std::mutex> lock(m_mutex);
                m_isStarted = false;
                m_cond.notify_all();
            }

            for (Threads::iterator it = m_threads.begin(); it != m_threads.end() ; ++it) {
                (*it)->join();
                delete *it;

                cout << "join thread" << endl;
            }
            m_threads.clear();
        }
    }

    void addTask(const Task& task, taskPriorityE level = level2) {
        // 2. 操作任务队列时要加锁
        std::unique_lock<std::mutex> lock(m_mutex);

        TaskPair taskPair(level, task);
        m_tasks.push(taskPair);

        // 3. 使用notify_one避免虚假唤醒
        m_cond.notify_one();
    }

private:
    void threadLoop() {
        while(m_isStarted)
        {
            Task task = take();
            if(task) {
                task();
            }
        }
    }
    Task take() {
        std::unique_lock<std::mutex> lock(m_mutex);

        //3. 使用 while-loop，避免虚假唤醒
        while(m_tasks.empty() && m_isStarted)
        {
            cout << "thread wait" << endl;
            m_cond.wait(lock);
        }

        Task task;
        Tasks::size_type size = m_tasks.size();
        if(!m_tasks.empty() && m_isStarted)
        {
            task = m_tasks.top().second;
            m_tasks.pop();
            assert(size - 1 == m_tasks.size());
        }

        return task;
    }

private:
    ThreadPool(const ThreadPool&);      //禁止复制拷贝
    const ThreadPool& operator=(const ThreadPool&);

    struct TaskPriorityCmp
    {
        bool operator()(const ThreadPool::TaskPair p1, const ThreadPool::TaskPair p2)
        {
            return p1.first > p2.first; //first的小值优先
        }
    };


private:
    typedef std::vector<std::thread*> Threads;
    typedef std::priority_queue<TaskPair, std::vector<TaskPair>, TaskPriorityCmp> Tasks;

    Threads m_threads;
    Tasks m_tasks;

    std::mutex m_mutex;
    condition_variable m_cond;  //利用条件变量，生产者唤醒消费者
    bool m_isStarted;
};

void sleep3() {
    cout << "start sleep 3" << endl;

    sleep(3);

    cout << "end sleep 3" << endl;
}

int main() {
    ThreadPool pool;

    pool.addTask(sleep3);
    pool.addTask(sleep3);
    pool.addTask(sleep3);
    pool.addTask(sleep3);

    sleep(10);
    return 0;
}