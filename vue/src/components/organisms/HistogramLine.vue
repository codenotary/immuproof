<template>
	<div class="histogram-line d-flex justify-space-between align-center full-width col-12">
		<div>{{ firstFormatted }}</div>
		<span class="histogram-line__line"></span>
		<div>{{ lastFormatted }}</div>
	</div>
</template>

<script>
import { formattedDateLocaleString } from '@/helpers/helpers';

export default {
	props: {
		last: { type: String, required: true },
		first: { type: String, required: true },
	},
	computed: {
		lastFormatted() {
			return formattedDateLocaleString(this.last, {
				month: 'short',
				timeZoneName: 'short',
			});
		},
		firstFormatted() {
			return formattedDateLocaleString(this.first, {
				month: 'short',
				timeZoneName: 'short',
			});
		},
	},
};
</script>
