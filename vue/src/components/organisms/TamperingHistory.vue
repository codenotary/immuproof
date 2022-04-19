<template>
	<div class="tampering-history full-width mt-10">
		<v-container class="tampering-history-wrapper">
			<v-row>
				<div class="tampering-history-wrapper__texts col-12">
					<div class="text-left font-weight-bold">State Check History</div>
				</div>
				<div class="tampering-history-wrapper__check col-12 flex justify-end mt-3">
					<span
						v-for="(item, index) in historyData"
						:key="index"
						class="tampering-history-wrapper__checks"
						:class="getClass(item)"
						@mouseover="toggleBox(index, true)"
						@mouseleave="toggleBox(index, false)"
					>
						<hover-box
							v-show="historyData[index].show"
							title="Proof Value"
							:history-data="item"
						>
						</hover-box>
					</span>
				</div>
				<histogram-line
					class="mt-3"
					:first="utcCheckFirst"
					:last="utcCheckLast"
				>
				</histogram-line>
			</v-row>
		</v-container>
	</div>
</template>

<script>
import HoverBox from '@/components/organisms/HoverBox.vue';
import HistogramLine from '@/components/organisms/HistogramLine.vue';

export default {
	components: { HistogramLine, HoverBox },
	props: {
		historyData: { type: Array, default: () => ([]) },
		utcCheckFirst: { type: String, default: '' },
		utcCheckLast: { type: String, default: '' },
	},
	methods: {
		getClass({ status }) {
			switch (status) {
				case 'UNKNOWN':
					return 'tampering-history-wrapper__checks-unknown';
				case 'CORRUPTED_DATA':
					return 'tampering-history-wrapper__checks-error';
				default:
					return 'tampering-history-wrapper__checks-normal';
			}
		},
		toggleBox(index, show) {
			this.$set(this.historyData[index], 'show', show);
		},
	},
};
</script>
