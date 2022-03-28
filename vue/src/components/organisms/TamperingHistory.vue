<template>
    <div class="tampering-history full-width mt-10">
        <v-container class="tampering-history-wrapper">
            <v-row>
                <div class="tampering-history-wrapper__texts col-12">
                    <div class="text-left font-weight-bold">Tampering Checks History</div>
                    <div class="text-left">{{ checkDate }}</div>
                </div>
                <div class="tampering-history-wrapper__check col-12 flex justify-center mt-3">
                    <span
                        v-for="(item, index) in data"
                        :key="index"
                        :data-information="getInformation(item)"
                        class="tampering-history-wrapper__checks"
                        :class="getClass(item)">
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
export default {
    props: {
        checkDate: {
            type: String,
            required: true
        },
        data: {
            type: Array,
            default: () => ([])
        }
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
        getInformation(item) {
            const date = new Date(item.time);
            const timeDetail = `${date.toDateString()} at ${date.toTimeString().split(' ')[0]}`

            return `${timeDetail} ${item.status}`;
        }
    }
}
</script>