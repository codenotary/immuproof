export function formattedDateLocaleString(date, extraOptions) {
	const options = { month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit', ...extraOptions };

	return new Date(date).toLocaleTimeString(undefined, options);
}