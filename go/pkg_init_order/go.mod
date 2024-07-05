module pkg_init_order_main

go 1.17

replace (
	github.com/miscmo/dotcode/go/pkg1 => ../pkg1
	github.com/miscmo/dotcode/go/pkg2 => ../pkg2
)

require (
	github.com/miscmo/dotcode/go/pkg1 v0.0.0-00010101000000-000000000000 // indirect
	github.com/miscmo/dotcode/go/pkg2 v0.0.0-00010101000000-000000000000 // indirect
)
