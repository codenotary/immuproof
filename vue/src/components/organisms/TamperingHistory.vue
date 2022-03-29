<template>
    <div class="tampering-history full-width mt-10">
        <v-container class="tampering-history-wrapper">
            <v-row>
                <div class="tampering-history-wrapper__texts col-12">
                    <div class="text-left font-weight-bold">Status Checks History</div>
                    <div class="text-left">{{ lastCheckDate }}</div>
                </div>
                <div class="tampering-history-wrapper__check col-12 flex justify-center mt-3">
                    <span
                        v-for="(item, index) in historyData"
                        :key="index"
                        class="tampering-history-wrapper__checks"
                        :class="getClass(item)"
                        @mouseover="showBox(index)"
                        @mouseleave="hideBox(index)">
                        <hover-box
                            v-show="historyData[index].show"
                            title="Proof Value"
                            :history-data="item">
                        </hover-box>
                    </span>
                </div>
                <div class="tampering-history-wrapper__last-check d-flex justify-space-between align-center full-width col-12 mt-3">
                    <div>Last check</div>
                    <span class="tampering-history-wrapper__last-check-line"></span>
                    <div>Today</div>
                </div>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import HoverBox from "@/components/organisms/HoverBox";
export default {
    components: { HoverBox },
    props: {
        lastCheckDate: {
            type: String,
            required: true
        },
        historyData: {
            type: Array,
            default: () => ([])
        }
    },

    data() {
       return {
           boxShow: []
       };
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
        showBox(index) {
            this.$set(this.historyData[index], 'show', true);
        },
        hideBox(index) {
            this.$set(this.historyData[index], 'show', false);
        }
    }
}
</script>