<template>
    <v-btn
        :href="link"
        class="d-flex align-center overflow-hidden mt-2"
        target="_blank"
        rel="noopener"
        :depressed="depressed"
        :outlined="outlined"
        :color="outlined ? color : ''">
        <slot
            v-if="icon !== ''"
            name="icon">
            <v-img
                :src="require(`@/assets/images/${icon}.svg`)"
                :contain="true"
                height="20">
            </v-img>
        </slot>
        <v-divider
            v-if="divider"
            class="my-4"
            style="max-height: 20px; height: 20px;"
            :color="color"
            :vertical="true">
        </v-divider>
        <slot></slot>
    </v-btn>
</template>

<script>
export default {
    props: {
        link: {
            type: String,
            default: ''
        },
        outlined: {
            type: Boolean,
            default: true
        },
        depressed: {
            type: Boolean,
            default: true
        },
        divider: {
            type: Boolean,
            default: false
        },
        color: {
            type: String,
            default: '#ffffff'
        },
        icon: {
            type: String,
            default: ''
        }
    }
}
</script>