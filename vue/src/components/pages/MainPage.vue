<template>
    <main-page
        tampering-message="NO TAMPERING DETECTED"
        check-date="Monday, 12 January 2022 at 15:34">
    </main-page>
</template>

<script>
import MainPage from '../templates/MainPage';

export default {
    components: {
        MainPage
    },

    beforeMount() {
        fetch('/api/status')
            .then(response => response.json())
            .then(data => console.log(data));
    }
}
</script>
