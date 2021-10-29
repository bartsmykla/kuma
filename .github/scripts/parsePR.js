module.exports = async ({
  github = {},
  owner = 'kumahq',
  kumaRepoName = 'kuma',
}, foundPR = null) => {
  if (!foundPR) {
    return null
  }

  try {
    const { number, title } = foundPR;

    const fullPR = await github.rest.pulls.get({
      owner,
      repo: kumaRepoName,
      pull_number: number,
    });

    const { data = {} } = fullPR;

    const {
      user = {},
      merged_by: mergedBy = {},
      labels: fullLabels = [],
    } = data;

    const labels = fullLabels.reduce((acc, { id, name, description }) => [
      ...acc,
      { id, name, description },
    ], []);

    console.log("fullPR", fullPR);
    console.log("labels", labels);

    return {
      number: number,
      title: title,
      openedBy: {
        login: user.login,
        id: user.id,
      },
      mergedBy: {
        login: mergedBy.login,
        id: mergedBy.id,
      },
      labels,
    };
  } catch (e) {
    console.error(e);
  }

  return null;
};
