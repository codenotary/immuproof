<template>
    <div class="floating-menu">
        <v-img
            class="floating-menu__logo"
            :src="require('@/assets/images/logo-cas.svg')"
            :contain="true"
        />
        <v-img
            class="floating-menu__cas-text"
            :src="require('@/assets/images/logo-cas-text.svg')"
            :contain="true"
            height="50"
        />
        <div class="floating-menu__buttons">
            <documentation-button></documentation-button>
            <discord-button></discord-button>
            <main-button
                icon="mdi-comment-quote"
                link="https://github.com/codenotary/immuproof/issues"
                color="#153954"
                :outlined="true"
                :divider="true">
                <div class="ml-2 text-capitalize">Share feedback</div>
            </main-button>
        </div>
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
