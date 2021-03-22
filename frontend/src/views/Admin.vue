<template>
  <v-app>
    <v-main>
      <KnowledgeGraph :width="1.0" :height="0.8" v-if="dataReady" :graph="this.g" v-on:concept-clicked="selectConcept">
        <!-- Modal popup button for concept creation -->
        <v-btn
          fab
          dark
          absolute
          color="primary"
          style="bottom: 32px; right: 32px;"
          @click="createNew"
          >
          <v-icon dark>
            mdi-plus
          </v-icon>
        </v-btn>
      </KnowledgeGraph>
      <MarkdownEditor 
        v-if="selectedConcept != null" 
        :title="selectedConcept.title" 
        :content="selectedConcept.content" 
        v-on:update:content="updateContent" 
        v-on:save="saveContent" 
        v-on:delete="deleteConcept"
      />
    
      <v-divider class="mx-8"></v-divider>

      <NewConceptModal v-if="showModal" v-on:close-modal="closeModal"/>
      <!-- Loading indicator -->
      <v-overlay v-if="!dataReady">
        <v-progress-circular
          indeterminate
          size="64"
        ></v-progress-circular>
      </v-overlay>
    </v-main>
  </v-app>
</template>

<script>
import KnowledgeGraph from '../components/KnowledgeGraph.vue'
import MarkdownEditor from '../components/MarkdownEditor.vue';
import NewConceptModal from '../components/NewConceptModal.vue';

import { getAllConcepts, updateConcept, deleteConcept } from '../services/ConceptService'

import * as dagreD3 from 'dagre-d3'

export default {
  name: 'Admin',
  components: {
    KnowledgeGraph,
    MarkdownEditor,
    NewConceptModal,
  },
  data() {
    return {
      g: null,
      dataReady: false,
      selectedConcept: null,
      showModal: false,
    }
  },
  computed: {
  },
  async mounted() {
    await this.getConcepts()
  },
  methods: {
    createNew() {
      this.showModal = true
    },
    closeModal() {
      this.showModal = false
      this.getConcepts()
    },
    updateContent(newContent) {
      this.selectedConcept.content = newContent
    },
    async getConcepts() {
      this.dataReady = false
      
      const res = await getAllConcepts()

      let graph = new dagreD3.graphlib.Graph()
      graph.setGraph({rankdir: "LR"})
      graph.setDefaultEdgeLabel(function() { return {} })

      if(res.concepts != null) {
        for(let i=0; i < res.concepts.length; i++) {
          graph.setNode(res.concepts[i].uuid, { label: res.concepts[i].title, shape: "circle", content: res.concepts[i].content }) 
        }
      }
      
      if(res.links != null) {
        for(let i=0; i < res.links.length; i++) {
          graph.setEdge(res.links[i].start_id, res.links[i].end_id) 
        }
      }

      this.g = graph
      this.dataReady = true
    },
    selectConcept(uuid) {
      const node = this.g.node(uuid)
      this.selectedConcept = {
        "uuid": uuid,
        "title": node.label,
        "content": node.content
      }
    },
    async saveContent() {
      // Run the update
      await updateConcept(this.selectedConcept.uuid, this.selectedConcept.content) 
      await this.getConcepts() 
    },
    async deleteConcept() {
      // Run the update
      await deleteConcept(this.selectedConcept.uuid)
      await this.getConcepts()
      this.selectedConcept = null 
    }
  },
}
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:wght@300;400;500&family=Roboto:wght@300;400;500&display=swap');

$body-font-family: 'Roboto';
$title-font: 'IBM Plex Sans';

.v-application {
  font-family: $body-font-family, sans-serif !important;
  .text-h1, .text-h2, .text-h3 { // To pin point specific classes of some components
       font-family: $title-font, sans-serif !important;
  }
}
</style>
