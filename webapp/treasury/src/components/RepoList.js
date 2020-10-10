import React from "react";
import {
  Avatar,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Paper,
  Typography,
} from "@material-ui/core";

import FolderIcon from "@material-ui/icons/Folder";
import { makeStyles } from "@material-ui/styles";
import { Link } from "react-router-dom";

export function RepoItem({ name, lastUpdated }) {
  return (
    <Link to={"/repos/" + name} key={name}>
      <ListItem button dense>
        <ListItemAvatar>
          <Avatar>
            <FolderIcon />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary={name} secondary={lastUpdated}></ListItemText>
      </ListItem>
    </Link>
  );
}

const useStyles = makeStyles(() => ({
  paper: {
    padding: 10,
    margin: 10,
    flex: 1,
    display: "flex",
    flexDirection: "column",
  },
  title: {
    marginBottom: 5,
  },
  wrapper: {
    flex: 1,
    overflow: "auto",
  },
}));

/*
Data is array of {name, lastUpdated}
*/
export function RepoList({ data }) {
  const styles = useStyles();

  function generateRepoList(data) {
    return data.map((element) => (
      <div key={element.name}>{RepoItem(element)}</div>
    ));
  }

  return (
    <Paper className={styles.paper}>
      <Typography variant="h4" className={styles.title}>
        {data ? "Projects" : "No repositories to show"}
      </Typography>
      <div className={styles.wrapper}>
        <List>{generateRepoList(data)}</List>
      </div>
    </Paper>
  );
}
