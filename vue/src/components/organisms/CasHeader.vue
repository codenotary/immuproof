<template>
    <div class="cas-header elevation-6 text-center full-width pa-8">
        <h2 class="text-uppercase font-weight-bold">Community Attestation Service Status Page</h2>
        <v-img
            class="cas-header__stamp"
            :src="require('@/assets/images/img-stamp.svg')"
            height="auto"
            :contain="true"
        />
        <div class="cas-header__hosted-by d-flex justify-center align-center mt-4">
            <p class="cas-header__hosted-by-text ma-0 mr-2 font-weight-bold">Hosted by:</p>
            <a
                href="https://almalinux.org/"
                target="_blank">
                <v-img
                    :src="require('@/assets/images/logo-home-assistant.svg')"
                    :max-height="30"
                    :max-width="140"
                    :contain="true"
                    height="auto">
                </v-img>
            </a>
            <a
                href="https://almalinux.org/"
                class="ma-0 ml-2 text-decoration-none">
                Home Assistant
            </a>
        </div>
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