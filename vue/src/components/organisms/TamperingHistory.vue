<template>
    <div class="tampering-history full-width mt-10">
        <v-container class="tampering-history-wrapper">
            <v-row>
                <div class="tampering-history-wrapper__texts col-12">
                    <div class="text-left font-weight-bold">Tampering Checks History</div>
                    <div class="text-left">Since Monday, 12 January 2022 at 15:34</div>
                </div>
                <div class="tampering-history-wrapper__check col-12 flex justify-center mt-3">
                    <span
                        v-for="index in 30"
                        :key="index"
                        :data-information="index"
                        class="tampering-history-wrapper__checks">
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

}
</script>

<style scoped>

</style>