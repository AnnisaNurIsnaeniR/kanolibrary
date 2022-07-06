<template>
  <v-row>
    <v-col cols="12" class="text-right">

      <v-btn
        depressed
        color="primary"
        :to="'/books/id'"
      >
        Add Book
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </v-col>
    <v-col cols="12">
      <v-data-table
        :headers="headers"
        :items="dataBuku"
        :items-per-page="5"
        class="elevation-1"
      >
        <template v-slot:[`item.actions`]="{ item }">
          <v-btn class="mx-2" fab dark small color="primary" @click="editBook(item)">
            <v-icon>mdi-pencil</v-icon>
          </v-btn>
          <v-btn class="mx-2" fab dark small color="warning" @click="deleteBook(item)">
            <v-icon>mdi-trash-can</v-icon>
          </v-btn>
        </template>
      </v-data-table>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'InspirePage',
  data() {
    return {
      headers: [
        {
          text: 'Title',
          align: 'start',
          sortable: false,
          value: 'Title',
        },
        { text: 'Author', value: 'Author' },
        { text: 'Publisher', value: 'Publisher' },
        { text: 'Synopsis', value: 'Synopsis' },
        { text: '', value: 'actions' },
      ],
      dataBuku: [],
      loading: false
    }
  },
  mounted(){
    this.getData()

  },
  methods: {
    editBook(item) {
      this.$router.push({
        path: `/Books/${item.Id}`,
      })
    },
    deleteBook(item){
      this.loading = true
          this.$axios.$delete('/api/books/delete?id=' + item.Id)
    .then(response =>{
      this.dataBuku = this.dataBuku.filter(data => data.id !== item.Id)
      this.getData()
      this.loading = false
    })   }
    ,  getData() {
    this.loading = true
    this.$axios.$get('/api/books')
    .then(response =>{
      this.dataBuku = response
      this.loading = false
    })
  }
  },

}
</script>
