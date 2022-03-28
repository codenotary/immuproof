export function formattedDateLocaleString(
	date,
	extraOptions = {
		weekday: 'long', year: 'numeric', month: 'long', day: 'numeric'
	}
) {
	const options = { ...extraOptions, weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };

	return new Date(date).toLocaleTimeString('en-US', options);
}

export function numFormatter(num) {
	const suffixes = ["", "k", "m", "b", "t", ''];
	const suffixNum = Math.floor((""+num).length/3);
	let shortValue = parseFloat((suffixNum !== 0 ? (num / Math.pow(1000,suffixNum)) : num).toPrecision(2));
	if (shortValue % 1 !== 0) {
		shortValue = shortValue.toFixed(1);
	}
	return shortValue+suffixes[suffixNum];
}