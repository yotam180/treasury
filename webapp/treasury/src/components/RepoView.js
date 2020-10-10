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
import React from "react";
import { withRouter } from "react-router";

import logo from "../assets/logo.png";

const useStyles = makeStyles(() => ({
  container: {
    margin: 10,
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
          <div className={styles.wrapper}>
            {ReleaseItem({ version: "0.1.3", lastUpdated: "yesterday" })}
            {ReleaseItem({ version: "0.1.2-alpha", lastUpdated: "2 days ago" })}

            {ReleaseItem({ version: "0.1.3", lastUpdated: "yesterday" })}
            {ReleaseItem({ version: "0.1.2-alpha", lastUpdated: "2 days ago" })}

            {ReleaseItem({ version: "0.1.3", lastUpdated: "yesterday" })}
            {ReleaseItem({ version: "0.1.2-alpha", lastUpdated: "2 days ago" })}

            {ReleaseItem({ version: "0.1.3", lastUpdated: "yesterday" })}
            {ReleaseItem({ version: "0.1.2-alpha", lastUpdated: "2 days ago" })}

            {ReleaseItem({ version: "0.1.3", lastUpdated: "yesterday" })}
            {ReleaseItem({ version: "0.1.2-alpha", lastUpdated: "2 days ago" })}

            {ReleaseItem({ version: "0.1.3", lastUpdated: "yesterday" })}
            {ReleaseItem({ version: "0.1.2-alpha", lastUpdated: "2 days ago" })}

            {ReleaseItem({ version: "0.1.3", lastUpdated: "yesterday" })}
            {ReleaseItem({ version: "0.1.2-alpha", lastUpdated: "2 days ago" })}
          </div>
        </Card>
      </Grid>
      <Grid item xs={9} className={styles.grid}>
        Hello world
      </Grid>
    </Grid>
  );
});
