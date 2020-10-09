import { Container, makeStyles } from "@material-ui/core";
import React from "react";
import { RepoList } from "../components/RepoList";

const useStyles = makeStyles(() => ({
  container: {
    height: "100%",
    display: "flex",
  },
}));

export function ReposView() {
  const styles = useStyles();

  function generateData() {
    let data = [];
    for (var i = 0; i < 50; ++i) {
      data.push({ name: "A" + Math.random(), lastUpdated: "yesterday" });
    }
    return data;
  }

  return (
    <Container className={styles.container}>
      <RepoList data={generateData()} />
    </Container>
  );
}
