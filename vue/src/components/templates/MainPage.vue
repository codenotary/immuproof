<template>
    <div class="main-page">
        <floating-menu></floating-menu>
        <v-container>
            <v-row class="text-center">
                <cas-header></cas-header>
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
                        :notarization-count-categories="notarizationCountCategories"
                        :notarization-count-data="notarizationCountData">
                    </notarizations>
                </div>
            </v-row>
        </v-container>
        <cas-footer></cas-footer>
    </div>
</template>

<script>
import FloatingMenu from '@/components/organisms/FloatingMenu';
import CasHeader from '@/components/organisms/CasHeader';
import TamperingStatus from '@/components/organisms/TamperingStatus';
import TamperingHistory from '@/components/organisms/TamperingHistory';
import Notarizations from '@/components/organisms/Notarizations';
import CasFooter from '@/components/organisms/CasFooter';

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
        tamperingMessage: {
            type: String,
            required: true
        },
        lastCheckDate: {
            type: String,
            required: true
        },
        firstCheckDate: {
            type: String,
            required: true
        },
        data: {
            type: Array,
            default: () => ([])
        },
        notarizations: {
            type: Array,
            default: () => ([])
        },
        notarizationCountCategories: {
            type: Array,
            default: () => ([])
        },
        notarizationCountData: {
            type: Array,
            default: () => ([])
        },
        lastTxId: {
            type: String,
            default: ''
        }
    },
}
</script>