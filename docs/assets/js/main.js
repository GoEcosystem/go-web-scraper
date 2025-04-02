/**
 * GoEcosystem Documentation - Main JavaScript
 * This file contains shared interactive features for the documentation site
 */

document.addEventListener('DOMContentLoaded', function() {
  // Mobile navigation toggle
  const mobileNavToggle = document.querySelector('.mobile-nav-toggle');
  if (mobileNavToggle) {
    mobileNavToggle.addEventListener('click', function() {
      document.querySelector('.site-nav').classList.toggle('active');
      this.setAttribute('aria-expanded', 
        this.getAttribute('aria-expanded') === 'false' ? 'true' : 'false'
      );
    });
  }

  // Add copy buttons to code blocks
  const codeBlocks = document.querySelectorAll('pre');
  codeBlocks.forEach(function(codeBlock) {
    const copyButton = document.createElement('button');
    copyButton.className = 'copy-button';
    copyButton.textContent = 'Copy';
    
    codeBlock.appendChild(copyButton);
    
    copyButton.addEventListener('click', function() {
      const code = codeBlock.querySelector('code') || codeBlock;
      const textToCopy = code.textContent;
      
      navigator.clipboard.writeText(textToCopy).then(function() {
        copyButton.textContent = 'Copied!';
        setTimeout(function() {
          copyButton.textContent = 'Copy';
        }, 2000);
      }, function() {
        copyButton.textContent = 'Failed to copy';
        setTimeout(function() {
          copyButton.textContent = 'Copy';
        }, 2000);
      });
    });
    
    // Add hover class to pre when hovering over it or the copy button
    codeBlock.addEventListener('mouseenter', function() {
      this.classList.add('hovered');
    });
    
    codeBlock.addEventListener('mouseleave', function() {
      this.classList.remove('hovered');
    });
  });
  
  // Add anchor links to headings
  const headings = document.querySelectorAll('h2[id], h3[id], h4[id], h5[id], h6[id]');
  headings.forEach(function(heading) {
    const anchor = document.createElement('a');
    anchor.className = 'heading-anchor';
    anchor.href = '#' + heading.id;
    anchor.innerHTML = '<span class="anchor-icon">#</span>';
    heading.appendChild(anchor);
  });
  
  // Handle active state in navigation based on current page
  const currentPath = window.location.pathname;
  const navLinks = document.querySelectorAll('.site-nav a');
  
  navLinks.forEach(function(link) {
    const linkPath = link.getAttribute('href');
    
    // Check if current path starts with the link path (for nested pages)
    if (currentPath === linkPath || 
        (linkPath !== '/' && currentPath.startsWith(linkPath))) {
      link.classList.add('active');
    }
  });
  
  // Initialize search if present
  const searchInput = document.getElementById('search-input');
  if (searchInput) {
    initSearch(searchInput);
  }
});

// Search functionality
function initSearch(searchInput) {
  const searchResults = document.getElementById('search-results');
  
  searchInput.addEventListener('input', function() {
    const query = this.value.trim().toLowerCase();
    
    if (query.length < 2) {
      searchResults.style.display = 'none';
      return;
    }
    
    // This would normally fetch from search index - simplified for demo
    fetch('/search-index.json')
      .then(response => response.json())
      .then(searchIndex => {
        const results = searchIndex.filter(item => 
          item.title.toLowerCase().includes(query) || 
          item.content.toLowerCase().includes(query)
        ).slice(0, 10);
        
        renderSearchResults(results, query, searchResults);
      })
      .catch(error => {
        console.error('Error performing search:', error);
        searchResults.innerHTML = '<p>Search is currently unavailable</p>';
        searchResults.style.display = 'block';
      });
  });
  
  // Close search results when clicking outside
  document.addEventListener('click', function(event) {
    if (!searchInput.contains(event.target) && !searchResults.contains(event.target)) {
      searchResults.style.display = 'none';
    }
  });
}

function renderSearchResults(results, query, searchResults) {
  if (results.length === 0) {
    searchResults.innerHTML = '<p>No results found for "' + query + '"</p>';
  } else {
    let resultsHTML = '<ul>';
    results.forEach(result => {
      resultsHTML += `
        <li>
          <a href="${result.url}">
            <span class="result-title">${result.title}</span>
            <span class="result-snippet">${highlightMatch(result.snippet, query)}</span>
          </a>
        </li>
      `;
    });
    resultsHTML += '</ul>';
    searchResults.innerHTML = resultsHTML;
  }
  
  searchResults.style.display = 'block';
}

function highlightMatch(text, query) {
  const regex = new RegExp('(' + query + ')', 'gi');
  return text.replace(regex, '<mark>$1</mark>');
}
