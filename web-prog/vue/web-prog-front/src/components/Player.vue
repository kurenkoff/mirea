<template>
   <div>
       Name:{{player.id}}
       Position:{{player.position}}
       Nationality:{{player.nationality}}
       <input type="checkbox" @change="sendChange">
   </div>
</template>

<script>
    export default {
        name: 'Player',

        props: {
            id: String,
            player: Object
        },

        mounted() {
            this.player = this.$store.getters.onePlayer(this.id)
        },

        methods: {
          sendChange(){
              this.$store.dispatch('updateForDelete', this.id)
          }
        }
    }
</script>

<style scoped>

</style>