<template>
  <Navbar />
  <KnowledgeGraph v-if="dataReady" :concepts="this.concepts" v-on:concept-clicked="selectConcept"/>
  <MarkdownEditor v-if="selectedConcept != null" :title="selectedConcept.title" :content="selectedConcept.content" v-on:update:content="updateContent" v-on:save="saveContent"/>
</template>

<script>
import KnowledgeGraph from './components/KnowledgeGraph.vue'
import Navbar from './components/Navbar.vue'
import MarkdownEditor from './components/MarkdownEditor.vue';

export default {
  name: 'App',
  components: {
    KnowledgeGraph,
    Navbar,
    MarkdownEditor,
  },
  data() {
    return {
      baseUrl: 'http://localhost:3000',
      dataReady: false,
      concepts: [],
      selectedConcept: null
    }
  },
  computed: {
  },
  async mounted() {
    const newConcepts = await this.getConcepts()
    this.concepts = newConcepts
    this.dataReady = true
  },
  methods: {
    updateContent(newContent) {
      this.selectedConcept.content = newContent
    },
    async getConcepts() {
      // Go and retrieve the concepts
      const rawResponse = await fetch(`${this.baseUrl}/concept/all`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }})
      
      return await rawResponse.json()
    },
    selectConcept(uuid) {
      this.selectedConcept = this.concepts.find(elem => elem.uuid == uuid)
    },
    async saveContent() {
      // Run the update
      await fetch(`${this.baseUrl}/concept/update/${this.selectedConcept.uuid}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({content: this.selectedConcept.content})
      })
      
      const rawResponse = await fetch(`${this.baseUrl}/concept/all`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }})
      
      this.concepts = await rawResponse.json()
    }
  },
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:wght@400;500&family=Roboto:wght@300;400;500&family=Roboto+Mono&display=swap');

h1, h2, h3, h4, h5, h6 {
  font-family: 'IBM Plex Sans', sans-serif;
  color: #7938D8;
}

.primary-color {
  background-color: #7938D8;
}

#app {
  font-family: 'Roboto', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
