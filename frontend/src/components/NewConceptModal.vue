<template>
  <v-overlay>
    <v-card light min-width="40vw">
      <v-card-title>Create new concept</v-card-title>
      <v-container>
        <v-row>
          <v-text-field
            class="mx-4"
            label="Concept Name"
            v-model="conceptTitle"
            outlined
          ></v-text-field>
        </v-row>
        <v-row>
          <v-autocomplete
            class="mx-4"
            :items="prereqs"
            :loading="isLoading"
            :search-input.sync="search"
            v-model="select"
            label="Prerequisites"
            placeholder="Start typing to search"
            no-data-text="No prerequisites found"
            cache-items
            multiple
            chips
            outlined
            item-text="title"
            item-value="uuid"
            return-object
            ></v-autocomplete> 
        </v-row>
        <v-row class="mt-6">
          <v-spacer/>
          <v-btn
            class="mb-4 mr-4"
            text
            color="primary"
            @click="closeModal"
          >
            Close
          </v-btn> 
          <v-btn
            @click="createConcept"
            class="mb-4 mr-4"
            elevation="2"
            color="primary"
            :loading="creating"
          >
            Create
          </v-btn> 
        </v-row>
      </v-container>
    </v-card>
  </v-overlay>
</template>

<script>
export default {
  name: 'NewConceptModal',
  data() {
    return {
      isLoading: false,
      prereqs: [],
      baseUrl: 'http://localhost:3000',
      select: null,
      search: null,
      conceptTitle: "",
      creating: false,
    }
  },
  computed: {
    items () {
      return this.prereqs.map(prereq => {
        return prereq.title
      })
    }
  },
  methods: {
    closeModal() {
      this.$emit("close-modal")
    },
    async createConcept() {
      this.creating = true;

      const req = {
        title: this.conceptTitle,
        content: `# ${this.conceptTitle}`,
        prerequisites: (this.select == null) ? [] : this.select.map(elem => {
          return elem.uuid
        })
      }

      await fetch(`${this.baseUrl}/concept/new`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(req)
      })

      this.$emit("close-modal")
      this.creating = false
    },
  },
  watch: {
    async search (val) {
      // Items have already been requested
      if (this.isLoading) return

      this.isLoading = true

      const rawResponse = await fetch(`${this.baseUrl}/concept/search?` + new URLSearchParams({query: val}).toString(), {
        method: 'GET'})

      try {
        const res = await rawResponse.json()
        this.prereqs = (res.concepts == null) ? [] : res.concepts
        this.isLoading = false
      } catch(e) {
        console.log(e)
        // No response returns means nothing
        this.isLoading = false
        return
      }
    },
  }
}
</script>

<style scoped>

</style>
