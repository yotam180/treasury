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
              <Route path="/repo/:name">
                <div>This is a repo view</div>
              </Route>
            </Switch>
          </div>
        </div>
      </ThemeProvider>
    </Router>
  );
}

export default App;
