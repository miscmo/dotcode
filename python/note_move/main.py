from math import fabs
import os
import shutil
import sys
# coding=utf-8

# 笔记迁移：从Mybase迁移到Obsidian

HandleDir = ""
IgnorePath = ['.git', 'node_modules']

if len(sys.argv) != 2:
    print('参数有误')
    sys.exit(1)

workpath = sys.argv[1]
if not os.path.isdir(workpath):
    print('filepath error')
    sys.exit(1)

print('filepath: ' + workpath)

def HandleNote(filepath):
    if not os.path.isdir(filepath):
        sys.exit(1)

    for item in os.listdir(filepath):
        print("item: " + item, "filepath: " + filepath)

        basename = os.path.basename(filepath)
        file = os.path.join(filepath, item)
        if os.path.isfile(file) and item == "defnote.md":
            newFile = os.path.join(filepath, basename + ".md")

            print("mv " + file + " to " + newFile)

            #shutil.copyfile(defnoteFile, newFile)
            shutil.move(file, newFile)

        if os.path.isdir(file):
            HandleNote(file)

    #if not hasSub:
        #mdnote = filepath + ".md"
        #parentdir = filepath + "../"

def HandlePath(filepath):
    hasSub = False
    obsidianpath = filepath + "_ObsidianNote"
    if not os.path.exists(filepath):
        sys.exit(1)

    shutil.copytree(filepath, obsidianpath)

    HandleNote(obsidianpath)

if __name__ == "__main__":
    HandlePath(workpath)
