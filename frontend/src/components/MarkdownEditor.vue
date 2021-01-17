<template>
  <v-container class="md-editor">
    <v-row class="mt-4">
      <v-col cols="12">
        <h1 class="text-h4">{{ title }}</h1>
      </v-col>
    </v-row>
    <v-row class="mt-6 mb-4 px-3">
      <h1 class="text-h6">Concept Content</h1>
      <v-spacer />
      <v-btn depressed class="mx-2" @click="onDelete" color="primary">Delete</v-btn>
      <v-btn depressed @click="onSave" color="primary">Save</v-btn>
    </v-row>
    <v-row class="rounded-lg flex-nowrap md-wrapper" no-gutters>
      <v-col cols="6" class="">
        <textarea id="raw" :value="content" @input="update" class="grey lighten-4 pa-6" style="width: 100%; height: 100%; resize: none;"></textarea>
      </v-col>
      <v-divider vertical></v-divider>
      <v-col cols="6">
        <div id="processed" v-html="compiledMarkdown" class="pa-6 markdown-body" style="min-height: 80vh;"></div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import * as marked from 'marked'

export default {
  name: 'MarkdownEditor',
  emits: ["update:content", "save"],
  props: {
    title: String,
    content: String,
  },
  data () {
    return {
      baseUrl: 'http://localhost:3000',
    }
  },
  computed: {
    compiledMarkdown: function() {
      return marked(this.content);
    }
  },
  methods: {
    update(e) { 
      this.$emit("update:content", e.target.value);
    },
    onSave() {
      console.log("Saving...")
      this.$emit("save", this.input)
    },
    onDelete() {
      console.log("Deleting")
      this.$emit("delete")
    },
  }
}
</script>

<style>
#raw {
  font-family: 'Roboto Mono', monospace;
}

.md-wrapper {
  border-radius: 4px;
  border-width: thin;
  border-style: solid;
  border-color: rgba(0, 0, 0, 0.12);
}
</style>
