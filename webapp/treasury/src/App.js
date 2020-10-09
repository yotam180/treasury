import React from "react";
import "./style/App.css";
import "./style/Fonts.css";
import "@material-ui/core/styles/";

import {
  AppBar,
  Button,
  CssBaseline,
  ThemeProvider,
  Toolbar,
} from "@material-ui/core";

function App() {
  return (
    <div className="App">
      <ThemeProvider>
        <CssBaseline />
        <AppBar position="static">
          <Toolbar edge="start">hello world</Toolbar>
        </AppBar>
        <Button variant="contained">Hello world</Button>
        hello world
      </ThemeProvider>
    </div>
  );
}

export default App;
