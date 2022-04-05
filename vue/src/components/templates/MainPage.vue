<template>
    <div class="main-page">
        <floating-menu></floating-menu>
        <v-container>
            <v-row class="text-center">
                <cas-header
                    :hosted-by-text="hostedByText"
                    :title-text="titleText"
                    :logo-url="logoUrl">
                </cas-header>
                <div class="main-page__content full-width">
                    <tampering-status
                        :tampering-message="tamperingMessage"
                        :last-check-date="lastCheckDate"
                        :last-tx-id="lastTxId">
                    </tampering-status>
                    <tampering-history
                        :last-check-date="lastCheckDate"
                        :history-data="data">
                    </tampering-history>
                    <notarizations
                        :first-check-date="firstCheckDate"
                        :notarizations="notarizations"
                        :notarization-categories-count="notarizationCategoriesCount"
                        :notarization-count-data="notarizationCountData">
                    </notarizations>
                </div>
            </v-row>
        </v-container>
        <cas-footer></cas-footer>
    </div>
</template>

<script>
import FloatingMenu from '@/components/organisms/FloatingMenu.vue';
import CasHeader from '@/components/organisms/CasHeader.vue';
import TamperingStatus from '@/components/organisms/TamperingStatus.vue';
import TamperingHistory from '@/components/organisms/TamperingHistory.vue';
import Notarizations from '@/components/organisms/Notarizations.vue';
import CasFooter from '@/components/organisms/CasFooter.vue';

export default {
    components: {
        CasFooter,
        Notarizations,
        TamperingStatus,
        CasHeader,
        FloatingMenu,
        TamperingHistory
    },
    props: {
        tamperingMessage: { type: String, required: true },
        lastCheckDate: { type: String, required: true },
        firstCheckDate: { type: String, required: true },
        data: { type: Array, default: () => ([]) },
        notarizations: { type: Array, default: () => ([]) },
        notarizationCategoriesCount: { type: Array, default: () => ([]) },
        notarizationCountData: { type: Array, default: () => ([]) },
        lastTxId: { type: Number, default: 0 },
        logoUrl: { type: String, default: '' },
        hostedByText: { type: String, default: '' },
        titleText: { type: String, default: '' },
    }
}
</script>