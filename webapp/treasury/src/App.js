import React from "react";
import "./style/App.css";
import "./style/Fonts.css";
import "@material-ui/core/styles/";

import { CssBaseline, ThemeProvider } from "@material-ui/core";
import { Header } from "./components/Header";
import { theme } from "./style/theme";
import { Switch, Route, BrowserRouter as Router } from "react-router-dom";
import { makeStyles } from "@material-ui/styles";
import { ReposView } from "./views/ReposView";
import { RepoView } from "./views/RepoView";

const useStyles = makeStyles(() => ({
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
    alignItems: "stretch",
    flexDirection: "column",
    display: "flex",
  },
}));

function App() {
  const styles = useStyles();

  return (
    <Router>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <div className={styles.fixedWrapper}>
          <Header />
          <div className={styles.app}>
            <Switch>
              <Route path="/" exact>
                <ReposView />
              </Route>
              <Route path="/repos/:name/releases/:release">
                <RepoView />
              </Route>
              <Route path="/repos/:name">
                <RepoView />
              </Route>
            </Switch>
          </div>
        </div>
      </ThemeProvider>
    </Router>
  );
}

export default App;
