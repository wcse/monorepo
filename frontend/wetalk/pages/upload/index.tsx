import withLayout from 'hocs/withLayout';
import styles from './styles.module.scss';

const Upload = () => {
  return (
    <div className={styles.upload}>
      <div className={styles.upload_container}>
        <div className="mb-2 text-xl">
          Drag and drop your tracks & albums here
        </div>
        <label
          htmlFor="files"
          className="shadow bg-blue-500 hover:bg-blue-400 focus:outline-none text-white font-bold py-2 px-4 rounded cursor-pointer"
        >
          or choose files to upload
        </label>
        <input type="file" name="files" id="files" className="hidden" />
      </div>
    </div>
  );
};

export default withLayout(Upload);
