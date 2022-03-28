export function formattedDateLocaleString(date, extraOptions) {
	const options = { month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit', ...extraOptions };

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