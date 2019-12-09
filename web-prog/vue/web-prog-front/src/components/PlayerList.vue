<template>
    <div>
        <ul id="player-list">
            <li v-for="item in this.$store.getters.allPlayers" v-bind:key="item.id">
                <Player v-bind:id="item.id"/>
            </li>
        </ul>
        <button v-on:click="getPlayers">Get</button>
        <button v-on:click="deletePlayers">Delete</button>
    </div>
</template>

<script>
    import Player from "@/components/Player";
    export default {
        name: "PlayerList",
        components: {Player},

        methods: {
            getPlayers(){
                this.$store.dispatch('fetchPlayers')
            },
            deletePlayers(){
                this.$store.dispatch('deletePlayers')
            }
        }
    }
</script>

<style scoped>

</style>