<template>
    <main-button
        :outlined="true"
        :divider="true"
        link="https://discord.gg/4wuHaewsxp"
        icon="mdi-discord"
        color="#153954">
        <p class="text-capitalize ml-2">{{ count }} Members</p>
    </main-button>
</template>

<script>
import MainButton from '@/components/atoms/MainButton.vue';

export default {
    components: {
        MainButton
    },
    data() {
        return {
            count: null,
        };
    },
    async mounted() {
        try {
            const { data } = await this.$axios.get('https://discordapp.com/api/v9/invites/4wuHaewsxp?with_counts=true');

            if (data) {
                this.count = data.approximate_member_count || '';
            }
        }
        catch (err) {
            console.error(err);
        }
    },
};
</script>
