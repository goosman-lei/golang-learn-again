module tec-inf.com/ex005

go 1.17

replace tec-inf.com/understandinit => ./understand-init-func

require (
	tec-inf.com/foo1 v0.0.0-00010101000000-000000000000
	tec-inf.com/foo2 v0.0.0-00010101000000-000000000000
	tec-inf.com/understandinit v0.0.0-00010101000000-000000000000
)

replace tec-inf.com/foo1 => ./understand-embeded-foo1

replace tec-inf.com/foo2 => ./understand-embeded-foo2
