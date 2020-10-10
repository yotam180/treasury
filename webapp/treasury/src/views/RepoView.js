import {
  Avatar,
  Card,
  CardContent,
  CardMedia,
  Grid,
  ListItem,
  ListItemAvatar,
  ListItemSecondaryAction,
  ListItemText,
  makeStyles,
  Typography,
} from "@material-ui/core";
import Icon from "@material-ui/icons/Pages";
import React, { useEffect, useState } from "react";
import { withRouter } from "react-router";
import axios from "axios";

import { REMOTE_URL } from "../settings";
import { processDate } from "../common";
import logo from "../assets/logo.png";
import { Link } from "react-router-dom";

const useStyles = makeStyles(() => ({
  container: {
    padding: 10,
    flex: "1 1",
    overflow: "hidden",
  },
  card: {
    margin: 10,
    marginBottom: 5,
  },
  media: {
    height: 0,
    paddingTop: "50%", // 16:9
  },
  wrapper: {
    overflowY: "auto",
    minHeight: 0,
    height: "100%",
  },
  grid: {
    height: "100%",
    display: "flex",
    flexDirection: "column",
    flexShrink: 0,
  },
}));

function ReleaseItem({ version, lastUpdated, selected }, repoName) {
  return (
    <Link to={`/repos/${repoName}/releases/${version}`}>
      <ListItem button dense selected={selected}>
        <ListItemAvatar>
          <Avatar>
            <Icon />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary={version} secondary={lastUpdated}></ListItemText>
      </ListItem>
    </Link>
  );
}

window.axios = axios;

async function getReleases(repoName) {
  var response = await axios.get(REMOTE_URL + "/api/repos/" + repoName);
  if (response.status !== 200) {
    throw Error(response.statusText);
  }

  return response.data;
}

async function getRelease(repoName, releaseName) {
  var response = await axios.get(
    REMOTE_URL + "/api/repos/" + repoName + "/releases/" + releaseName
  );

  if (response.status !== 200) {
    throw Error(response.statusText);
  }

  return response.data;
}

/*
data = [{
        version: string
        last_updated: datetime-string
    }...
]
*/
export const RepoView = withRouter(function ({ match }) {
  const repoName = match.params.name;
  const releaseName = match.params.release;

  const styles = useStyles();

  const [state, setState] = useState({});
  const [error, setError] = useState(null);

  const [releaseState, setReleaseState] = useState({});

  useEffect(() => {
    getReleases(repoName).then(setState).catch(setError);
  }, [repoName]);

  useEffect(() => {
    if (releaseName) {
      getRelease(repoName, releaseName).then(setReleaseState).catch(setError);
    }
  }, [releaseName]);

  function generateReleaseList() {
    if (state.releases === undefined) {
      return <div>Loading...</div>;
    }

    if (state.releases.length === 0) {
      return <div>No releases to show...</div>;
    }

    return state.releases.map(({ version, last_updated }) => (
      <div key={version}>
        {ReleaseItem(
          {
            version,
            lastUpdated: processDate(last_updated),
            selected: version === releaseState?.data?.version,
          },
          repoName
        )}
      </div>
    ));
  }

  return (
    <Grid container className={styles.container}>
      <Grid item xs={3} className={styles.grid}>
        <Card className={styles.card}>
          <CardMedia
            image={logo}
            title="Repository name comes here"
            className={styles.media}
          />
          <CardContent>
            <Typography variant="h5">{repoName}</Typography>
          </CardContent>
        </Card>
        <Card className={styles.card} style={{ flex: 1 }}>
          <div className={styles.wrapper}>{generateReleaseList()}</div>
        </Card>
      </Grid>
      <Grid item xs={9} className={styles.grid}>
        <Card className={styles.card} style={{ flex: 1 }}>
          <div className={styles.wrapper}>
            {error != null && error.toString()}
            {releaseState && JSON.stringify(releaseState)}
          </div>
        </Card>
      </Grid>
    </Grid>
  );
});
