console.log('alp')

fetch('/api/status')
	.then(response => response.json())
	.then(data => console.log(data));