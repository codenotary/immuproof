export function formattedDateLocaleString(
	date,
	extraOptions = {
		weekday: 'long', year: 'numeric', month: 'long', day: 'numeric'
	}
) {
	console.log(date)
	const options = { ...extraOptions, weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };

	return new Date(date).toLocaleTimeString('en-US', options);
}