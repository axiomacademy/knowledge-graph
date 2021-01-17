import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: "#7938D8", // #E53935
        secondary: "#491496", // #FFCDD2
        accent: "#DD1155", // #3F51B5
      },
    },
  },
});
