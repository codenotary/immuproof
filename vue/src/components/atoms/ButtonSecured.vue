<template>
    <button-pill
        href="https://github.com/codenotary/immudb"
        class="immudb-button-secured-by"
    >
        <template #icon>
            <v-img
                class="mx-2"
                :src="require('@/assets/images/logo-only-small.svg')"
                alt="Immudb"
                :max-height="42"
                :max-width="140"
            />
        </template>
        <div class="dark--text caption pa-2">
            <div class="immudb-button-secured-by__text-small | line-height-1">
                Secured by
            </div>
            <div class="font-weight-bold line-height-1">
                immudb
            </div>
        </div>
    </button-pill>
</template>

<script>
import ButtonPill from '@/components/atoms/ButtonPill.vue';

export default {
    components: {
        ButtonPill
    }
}
</script>

<style lang="scss">
.immudb-button-secured-by {
    &__text-small {
        font-size: 7.5px;
    }
}
</style>
