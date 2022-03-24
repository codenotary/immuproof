<template>
    <div class="notarizations full-width mt-10">
        <v-container class="notarizations-wrapper">
            <v-row>
                <div class="col-12 notarizations-wrapper__texts">
                    <div class="text-left font-weight-bold">Notarizations</div>
                    <div class="text-left">Since Monday, 12 January 2022 at 15:34</div>
                </div>
                <chart class="mt-3"></chart>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import Chart from "@/components/organisms/Chart";
export default {
    components: {Chart}
}
</script>