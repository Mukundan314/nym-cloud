import React from "react";
import * as styles from "./index.module.css";
import propTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFile } from "@fortawesome/free-solid-svg-icons";
import classnames from "classnames";

const ListItem = ({ name, date, size, idx, selected, onSelect }) => (
  <div
    className={classnames(styles.listItem, {
      [styles.listItemSelected]: selected,
      [styles.listItemHighlighted]: idx % 2,
    })}
    onClick={onSelect}
    role="button"
  >
    <div className={styles.cell}>
      <FontAwesomeIcon icon={faFile} className={styles.itemIcon} />
      {name}
    </div>
    <div className={styles.cell}>{new Date(date).toLocaleString()}</div>
    <div className={styles.cell}>{size}</div>
  </div>
);

ListItem.propTypes = {
  name: propTypes.string.isRequired,
  date: propTypes.string.isRequired,
  size: propTypes.string.isRequired,
  idx: propTypes.number.isRequired,
  selected: propTypes.bool.isRequired,
  onSelect: propTypes.func.isRequired,
};

export default ListItem;
