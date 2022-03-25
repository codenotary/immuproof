console.log('alp')

fetch('/api/status')
	.then(response => response.json())
	.then(data => {
		console.log('data:', data);
		// var a = data.slice(1).slice(-30)
		// console.log('a:', a);
	});