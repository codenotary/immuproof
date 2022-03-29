<template>
    <div class="floating-menu d-flex flex-column">
        <v-img
            class="floating-menu__logo"
            :src="require('@/assets/images/logo-cas.svg')"
            :contain="true"
            height="90"
        />
        <v-img
            class="floating-menu__cas-text"
            :src="require('@/assets/images/logo-cas-text.svg')"
            :contain="true"
            height="50"
        />
        <documentation-button class="mt-3"></documentation-button>
        <discord-button></discord-button>
        <main-button
            icon="mdi-discord"
            link="https://github.com/codenotary/immuproof/issues"
            color="#153954"
            :outlined="true"
            :divider="true">
            <div class="ml-2 text-capitalize">Share feedback</div>
        </main-button>
    </div>
</template>

<script>
import DocumentationButton from '@/components/organisms/DocumentationButton.vue';
import MainButton from '@/components/atoms/MainButton.vue';
import DiscordButton from '@/components/organisms/DiscordButton.vue';

export default {
    components: {
        DiscordButton,
        MainButton,
        DocumentationButton
    }
}
</script>
