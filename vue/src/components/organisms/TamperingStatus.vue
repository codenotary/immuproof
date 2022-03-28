<template>
    <div class="tampering-status-wrapper full-width mt-10">
        <v-container>
            <v-row justify="space-between" class="tampering-status-wrapper__content align-center">
                <div
                    class="tampering-status-wrapper__status d-flex align-center"
                    :class="{ 'tampering-status-wrapper__status-error' : tamperingMessage !== 'NORMAL' }">
                    <p class="ma-0 ml-3">{{ tamperingMessage }}</p>
                </div>
                <div class="tampering-status-wrapper__last-check">
                    <p class="font-weight-bold text-left">Last Check</p>
                    <p class="tampering-status-wrapper__last-check-date text-left">{{ checkDate }}</p>
                    <p class="text-left">
                        <span class="font-weight-bold">TxID</span>
                        25
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
        checkDate: {
            type: String,
            required: true
        }
    }
}
</script>

<style scoped>

</style>