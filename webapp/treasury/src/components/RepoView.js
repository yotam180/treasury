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

function ReleaseItem({ version, lastUpdated }) {
  return (
    <ListItem button dense>
      <ListItemAvatar>
        <Avatar>
          <Icon />
        </Avatar>
      </ListItemAvatar>
      <ListItemText primary={version} secondary={lastUpdated}></ListItemText>
    </ListItem>
  );
}

window.axios = axios;

async function getReleases(repoName) {
  var response = await axios.get(REMOTE_URL + "/api/repos/" + repoName);
  if (response.status != 200) {
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
export const RepoView = withRouter(function ({ data, match }) {
  const releases = data;
  const repoName = match.params.name;

  const styles = useStyles();

  const [state, setState] = useState({});
  const [error, setError] = useState(null);

  useEffect(() => {
    console.log("requesting repo data");
    getReleases(repoName).then(setState).catch(setError);
  }, []);

  function generateReleaseList() {
    if (state.releases === undefined) {
      return <div>Loading...</div>;
    }

    if (state.releases.length === 0) {
      return <div>No releases to show...</div>;
    }

    return state.releases.map(({ version, last_updated }) => (
      <div key={version}>
        {ReleaseItem({ version, lastUpdated: processDate(last_updated) })}
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
          </div>
        </Card>
      </Grid>
    </Grid>
  );
});
