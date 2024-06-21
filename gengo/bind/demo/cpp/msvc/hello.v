module main
fn main() {
	println('hello,world')
	demo()
}

fn demo() {
	areas := ['game', 'web', 'tools', 'science', 'systems',
	'embedded', 'drivers', 'GUI', 'mobile']
	for area in areas {
		println('Hello, ${area} developers!')
	}
}