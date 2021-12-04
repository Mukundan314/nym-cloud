import React, { useState } from "react";
import * as styles from "./index.module.css";
import ListItem from "../ListItem";
import { fields } from "../../constants/listView";
import { files } from "../../constants/files";
import propTypes from "prop-types";

const ListView = ({ search }) => {
  const [selectedItem, setSelectedItem] = useState(null);
  return (
    <div className={styles.listView}>
      <div className={styles.fields}>
        {fields.map((field) => (
          <div className={styles.cell}>{field}</div>
        ))}
      </div>
      <div>
        {files
          .filter(({ name }) => name.includes(search))
          .map((file, idx) => (
            <div className={styles.listItem}>
              <ListItem
                name={file.name}
                date={file.date}
                size={file.size}
                idx={idx}
                selected={selectedItem === idx}
                onSelect={() => setSelectedItem(idx)}
              />
            </div>
          ))}
      </div>
    </div>
  );
};

ListView.propTypes = {
  search: propTypes.string.isRequired,
};

export default ListView;
