<template>
    <div class="tampering-status-wrapper full-width mt-10">
        <v-container>
            <v-row justify="space-between" class="tampering-status-wrapper__content align-center">
                <div
                    class="tampering-status-wrapper__status d-flex align-center"
                    :class="getClassName">
                    <p class="ma-0 ml-3">{{ tamperingMessage }}</p>
                </div>
                <div class="tampering-status-wrapper__last-check">
                    <p class="font-weight-bold text-left">Last Check</p>
                    <p class="tampering-status-wrapper__last-check-date text-left">{{ lastCheckDate }}</p>
                    <p class="tampering-status-wrapper__last-check-id text-left">
                        <span>TxID</span>
                        {{ lastTxId }}
                    </p>
                </div>
            </v-row>
        </v-container>
    </div>
</template>

<script>
export default {
    props: {
        tamperingMessage: {
            type: String,
            required: true
        },
        lastCheckDate: {
            type: String,
            required: true
        },
        lastTxId: {
            type: String,
            required: true
        }
    },

    computed: {
        getClassName() {
            switch (this.tamperingMessage) {
                case 'Status Unknown':
                    return 'tampering-status-wrapper__status-unknown';
                case 'Tampering Detected':
                    return 'tampering-status-wrapper__status-error';
                default:
                    return 'tampering-status-wrapper__status-normal';
            }
        }
    }
}
</script>

<style scoped>

</style>