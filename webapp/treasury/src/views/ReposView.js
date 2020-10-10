import { Container, makeStyles } from "@material-ui/core";
import axios from "axios";
import React, { useEffect, useState } from "react";
import { processDate } from "../common";
import { RepoList } from "../components/RepoList";
import { REMOTE_URL } from "../settings";

const useStyles = makeStyles(() => ({
  container: {
    height: "100%",
    display: "flex",
  },
}));

async function getRepos() {
  const response = await axios.get(REMOTE_URL + "/api/repos");
  if (response.status !== 200) {
    throw Error(response.statusText);
  }

  return response.data;
}

export function ReposView() {
  const styles = useStyles();

  const [state, setState] = useState([]);

  useEffect(() => {
    getRepos().then((response) =>
      setState(
        response.data.map(({ name, last_updated }) => ({
          name,
          lastUpdated: processDate(last_updated),
        }))
      )
    );
  }, []);

  function generateData() {
    let data = [];
    for (var i = 0; i < 50; ++i) {
      data.push({ name: "A" + Math.random(), lastUpdated: "yesterday" });
    }
    return data;
  }

  return (
    <Container className={styles.container}>
      <RepoList data={state || []} />
    </Container>
  );
}
