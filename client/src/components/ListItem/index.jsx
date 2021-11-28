import React from "react";
import * as styles from "./index.module.css";
import propTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFolder, faFile } from "@fortawesome/free-solid-svg-icons";
import classnames from "classnames";

const ListItem = ({ name, date, size, type, idx }) => (
  <div
    className={classnames(styles.listItem, {
      [styles.listItemHighlighted]: idx % 2,
    })}
  >
    <div>
      {type === "file" ? (
        <FontAwesomeIcon icon={faFile} className={styles.itemIcon} />
      ) : (
        <FontAwesomeIcon icon={faFolder} className={styles.itemIcon} />
      )}
      {name}
    </div>
    <div>{date}</div>
    <div>{size}</div>
    <div />
  </div>
);

ListItem.propTypes = {
  name: propTypes.string.isRequired,
  date: propTypes.string.isRequired,
  size: propTypes.string.isRequired,
  type: propTypes.string.isRequired,
  idx: propTypes.number.isRequired,
};

export default ListItem;
