import {
  Avatar,
  Card,
  CardContent,
  CardMedia,
  Chip,
  Grid,
  ListItem,
  ListItemAvatar,
  ListItemText,
  makeStyles,
  Paper,
  Typography,
} from "@material-ui/core";
import FileIcon from "@material-ui/icons/Description";
import Icon from "@material-ui/icons/Pages";
import axios from "axios";
import React, { useEffect, useState } from "react";
import { withRouter } from "react-router";
import { Link } from "react-router-dom";

import placeholder from "../assets/placeholder.png";
import { processDate } from "../common";
import { REMOTE_URL } from "../settings";

const useStyles = makeStyles((theme) => ({
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
  chipArray: {
    margin: 10,
    padding: 0,
    background: theme.palette.primary.main,
  },
  chip: {
    margin: 5,
    marginRight: 0,
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

// Taken from StackOverflow ;)
function downloadURI(uri, name) {
  var link = document.createElement("a");
  // If you don't know the name or want to use
  // the webserver default set name = ''
  link.setAttribute("download", name);
  link.href = uri;
  document.body.appendChild(link);
  link.click();
  link.remove();
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
  }, [releaseName, repoName]);

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

  function generateFilesList() {
    if (!releaseState) {
      return null;
    }

    if (!releaseState?.files?.length) {
      return <div>No files to show for release</div>;
    }

    return releaseState.files.map((file) => (
      <ListItem
        button
        dense
        key={file}
        onClick={() =>
          downloadURI(
            `${REMOTE_URL}/api/repos/${repoName}/releases/${releaseName}/files${file}`
          )
        }
      >
        <ListItemAvatar>
          <Avatar>
            <FileIcon />
          </Avatar>
        </ListItemAvatar>
        <ListItemText
          primaryTypographyProps={{
            variant: "body1",
          }}
          primary={file.substr(1)}
        ></ListItemText>
      </ListItem>
    ));
  }

  console.log(state);

  return (
    <Grid container className={styles.container}>
      <Grid item xs={3} className={styles.grid}>
        <Card className={styles.card}>
          <CardMedia
            image={state?.metadata?.picture || placeholder}
            title="Repository name comes here"
            className={styles.media}
          />
          <CardContent>
            <Typography variant="h5">{repoName}</Typography>
          </CardContent>
        </Card>
        <Card
          className={styles.card}
          style={{
            flex: 1,
          }}
        >
          <div className={styles.wrapper}>{generateReleaseList()}</div>
        </Card>
      </Grid>

      <Grid item xs={9} className={styles.grid}>
        {Object.keys(releaseState).length > 0 && (
          <Card className={styles.card} style={{ flex: 1, padding: 10 }}>
            {error != null && error.toString()}
            <Typography variant="h4">
              {repoName} / {releaseState?.data?.version}
            </Typography>
            <Paper component="ul" className={styles.chipArray}>
              {Object.keys(releaseState?.metadata || {}).map((key) => (
                <Chip
                  key={key}
                  className={styles.chip}
                  label={`${key}: ${releaseState?.metadata[key]}`}
                />
              ))}
            </Paper>

            <div className={styles.wrapper}>{generateFilesList()}</div>
          </Card>
        )}
      </Grid>
    </Grid>
  );
});
