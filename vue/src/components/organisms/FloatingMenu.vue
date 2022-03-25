<template>
    <div class="floating-container d-flex flex-column">
        <v-img
            class="menu-floating__logo"
            :src="require('@/assets/images/logo-cas.svg')"
            :contain="true"
            height="90"
        />
        <v-img
            class="menu-floating__logo"
            :src="require('@/assets/images/logo-cas-text.svg')"
            :contain="true"
            height="50"
        />
        <documentation-button class="mt-3"></documentation-button>
        <i-button
            icon="filled-discord"
            :divider="true"
            color="#147cb6">
            <div class="pa-2">Members</div>
        </i-button>
        <i-button
            :divider="true"
            icon="filled-discord"
            color="#147cb6">
            <div class="pa-2">Share feedback</div>
        </i-button>
    </div>
</template>

<script>
import DocumentationButton from '@/components/organisms/DocumentationButton.vue';
import IButton from '@/components/atoms/iButton.vue';

export default {
    components: {
        IButton,
        DocumentationButton
    }
}
</script>

<style scoped>
    .floating-container {
        position: fixed;
        top: 0;
        left: 0;
        padding: 25px;
    }
</style>