import React from "react";
import "./style/App.css";
import "./style/Fonts.css";
import "@material-ui/core/styles/";

import { Container, CssBaseline, ThemeProvider } from "@material-ui/core";
import { Header } from "./components/Header";
import { theme } from "./style/theme";
import { RepoList } from "./components/RepoList";
import { makeStyles } from "@material-ui/styles";

const useStyles = makeStyles((theme) => ({
  fixedWrapper: {
    top: 0,
    bottom: 0,
    left: 0,
    right: 0,
    display: "flex",
    position: "fixed",
    flexDirection: "column",
  },
  app: {
    flex: 1,
    overflow: "hidden",
  },
  container: {
    height: "100%",
    display: "flex",
  },
}));

function App() {
  function generateData() {
    let data = [];
    for (var i = 0; i < 50; ++i) {
      data.push({ name: "A" + Math.random(), lastUpdated: "yesterday" });
    }
    return data;
  }

  const styles = useStyles();

  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        <CssBaseline />

        <div className={styles.fixedWrapper}>
          <Header />
          <div className={styles.app}>
            <Container className={styles.container}>
              <RepoList data={generateData()} />
            </Container>
          </div>
        </div>
      </ThemeProvider>
    </div>
  );
}

export default App;
