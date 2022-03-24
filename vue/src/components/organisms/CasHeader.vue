<template>
    <div class="cas-header elevation-6 text-center full-width pa-8">
        <h2 class="text-uppercase font-weight-bold">Community Attestation Service Status Page</h2>
        <v-img
            class="cas-header__stamp"
            :src="require('@/assets/images/img-stamp.svg')"
            height="auto"
            :contain="true"
        />
    </div>
</template>

<script>
export default {}
</script>

<style scoped lang="scss">
    .cas-header {
        border-radius: 0 0 24px 24px;

        h2 {
            line-height: 1;
            color: #244583;
        }
    }
</style>