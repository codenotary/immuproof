<template>
    <v-app>
        <v-main>
            <main-page></main-page>
        </v-main>
    </v-app>
</template>

<script>
import MainPage from '@/components/pages/MainPage.vue';

export default {
    components: {
        MainPage
    },
    beforeCreate() {
        // Creates script tag that contains values to be populated by BE
        let serverScript = document.createElement('script');
        serverScript.innerHTML = 'window.portValue = "{{{ .PORT }}}"; window.address = "{{{ .ADDRESS }}}"; window.hostedByLogoUrl = "{{{ .HOSTED_BY_LOGO_URL }}}";'

        document.body.appendChild(serverScript);
    }
};
</script>

<style lang="scss">
@import "assets/styling/colors";

#app {
    background-color: $light-base;
}
</style>
