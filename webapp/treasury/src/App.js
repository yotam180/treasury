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
import { RepoView } from "./components/RepoView";

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
                <RepoView
                  data={[
                    {
                      version: "0.1.0-alpha",
                      last_updated: "2020-10-09T14:20:47.234204+03:00",
                    },
                    {
                      version: "0.1.1",
                      last_updated: "2020-10-08T23:40:47.2057345+03:00",
                    },
                    {
                      version: "0.2.3",
                      last_updated: "2020-10-08T23:40:52.7862415+03:00",
                    },
                    {
                      version: "0.3.0",
                      last_updated: "2020-10-09T11:15:38.5226762+03:00",
                    },
                  ]}
                />
              </Route>
            </Switch>
          </div>
        </div>
      </ThemeProvider>
    </Router>
  );
}

export default App;
