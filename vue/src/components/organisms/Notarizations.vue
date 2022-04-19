<template>
	<div class="notarizations full-width mt-10">
		<v-container class="notarizations-wrapper">
			<v-row>
				<div class="col-12 notarizations-wrapper__texts">
					<div class="text-left font-weight-bold">Notarizations</div>
				</div>
				<apex-chart
					class="mt-3"
					type="line"
					:categories="notarizationCategoriesCount"
					:data="notarizationCountData"
				>
				</apex-chart>
				<histogram-line
					class="notarizations__histogram-line"
					:first="utcCheckFirst"
					:last="utcCheckLast"
				>
				</histogram-line>
			</v-row>
		</v-container>
	</div>
</template>

<script>
import ApexChart from '@/components/organisms/ApexChart.vue';
import HistogramLine from '@/components/organisms/HistogramLine';

export default {
	components: {
		HistogramLine,
		ApexChart,
	},
	props: {
		notarizations: { type: Array, default: () => ([]) },
		notarizationCategoriesCount: { type: Array, default: () => ([]) },
		notarizationCountData: { type: Array, default: () => ([]) },
		utcCheckFirst: { type: String, default: '' },
		utcCheckLast: { type: String, default: '' },
	},
};
</script>