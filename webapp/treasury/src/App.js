import React from "react";
import "./style/App.css";
import "./style/Fonts.css";
import "@material-ui/core/styles/";

import { CssBaseline, ThemeProvider } from "@material-ui/core";
import { Header } from "./components/Header";
import { theme } from "./style/theme";

function App() {
  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Header />
        Hello, world!
      </ThemeProvider>
    </div>
  );
}

export default App;
