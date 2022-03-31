<template>
    <v-app>
        <v-main>
            <main-page></main-page>
        </v-main>
    </v-app>
</template>

<script>
import MainPage from "./components/pages/MainPage";

export default {
    components: {
        MainPage
    },
    beforeCreate() {
        let serverScript = document.createElement('script');
        serverScript.innerHTML = 'window.portValue = "{{{ .PORT }}}"; window.address = "{{{ .ADDRESS }}}"; window.hostedByLink = "{{{ .HOSTED_BY_LINK }}}";'

        document.body.appendChild(serverScript);
    }
};
</script>

<style lang="scss">
#app {
    background-color: #FAF9F8;
}
</style>
