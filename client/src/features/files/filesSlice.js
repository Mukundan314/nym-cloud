import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit";

const initialState = {
  files: [],
  status: "idle",
  error: null,
};

export const fetchFiles = createAsyncThunk("files/fetchFiles", async () => {
  // TODO
  await new Promise((resolve) => {
    setTimeout(resolve, 1000);
  });
  return {
    files: [
      {
        name: "Photos",
        type: "dir",
        size: 2,
        date: 1636898223,
        conent: [
          {
            name: "img1.png",
            date: 1636898223,
            size: "73KB",
            type: "file",
            hash: "hash",
          },
          {
            name: "img2.png",
            date: 1636898223,
            size: "73KB",
            type: "file",
            hash: "hash",
          },
        ],
      },
      {
        name: "hello.png",
        date: 1621000613,
        size: "73KB",
        type: "file",
        hash: "hash",
      },
    ],
  };
});

const filesSlice = createSlice({
  name: "files",
  initialState,
  reducers: {},
  extraReducers(builder) {
    builder
      .addCase(fetchFiles.pending, (state) => {
        state.status = "loading";
      })
      .addCase(fetchFiles.fulfilled, (state, action) => {
        state.status = "idle";
        state.files = action.files;
      })
      .addCase(fetchFiles.rejected, (state, action) => {
        state.status = "error";
        state.error = action.error;
      });
  },
});

export const selectFilesStatus = (state) => state.files.status;
export const selectFilesError = (state) => state.files.error;
export const selectAllFiles = (state) => state.files.files;

export default filesSlice.reducer;
