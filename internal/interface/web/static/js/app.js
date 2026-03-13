// Fighting Fantasy - Vanilla JavaScript utilities
// Simple, clean JavaScript for client-side interactions

/**
 * Tab switching utility
 * Used in auth and dashboard pages
 */
function switchTab(tabName) {
    const allTabs = document.querySelectorAll('.tab-content');
    const allButtons = document.querySelectorAll('.tab-button');
    
    // Hide all tab contents
    allTabs.forEach(tab => {
        tab.style.display = 'none';
    });
    
    // Remove active class from all buttons
    allButtons.forEach(btn => {
        btn.classList.remove('active');
    });
    
    // Show selected tab
    const selectedTab = document.getElementById(tabName + '-tab');
    if (selectedTab) {
        selectedTab.style.display = 'block';
    }
    
    // Add active class to clicked button
    if (event && event.target) {
        const button = event.target.closest('.tab-button');
        if (button) {
            button.classList.add('active');
        }
    }
}

/**
 * Potion selection in hero creation
 */
function selectPotion(potionType, element) {
    const allOptions = document.querySelectorAll('.potion-option');
    
    // Remove active styling from all options
    allOptions.forEach(option => {
        option.classList.remove('border-[var(--gold)]', 'bg-primary/20');
        option.classList.add('border-border');
    });
    
    // Add active styling to selected option
    element.classList.remove('border-border');
    element.classList.add('border-[var(--gold)]', 'bg-primary/20');
    
    // Update hidden input
    const hiddenInput = document.getElementById('selected-potion');
    if (hiddenInput) {
        hiddenInput.value = potionType;
    }
}

/**
 * Dice rolling functionality
 * @param {string} id - Dice identifier (player/enemy)
 * @param {number} sides - Number of sides on the dice (default 6)
 */
function rollDice(id, sides = 6) {
    const diceEl = document.getElementById('dice-' + id);
    const valueEl = document.getElementById('dice-value-' + id);
    const iconEl = document.getElementById('dice-icon-' + id);
    const buttonEl = document.getElementById('roll-button-' + id);
    
    if (!diceEl || !valueEl || !iconEl || !buttonEl) {
        console.error('Dice elements not found for id:', id);
        return;
    }
    
    // Disable button and update text
    buttonEl.disabled = true;
    buttonEl.textContent = 'Rolling...';
    
    // Hide value, show icon
    valueEl.style.display = 'none';
    iconEl.style.display = 'block';
    
    // Add rotation animation
    diceEl.style.transform = 'rotate(360deg)';
    
    // Simulate roll delay
    setTimeout(() => {
        // Generate random value
        const result = Math.floor(Math.random() * sides) + 1;
        
        // Show result
        valueEl.textContent = result;
        valueEl.style.display = 'block';
        iconEl.style.display = 'none';
        
        // Reset rotation
        diceEl.style.transform = 'rotate(0deg)';
        
        // Re-enable button
        buttonEl.disabled = false;
        buttonEl.textContent = 'Roll D' + sides;
        
        // Store result for battle calculations
        if (id === 'player') {
            window.lastPlayerRoll = result;
        } else if (id === 'enemy') {
            window.lastEnemyRoll = result;
        }
        
        // Fire callback if defined
        if (typeof window.onDiceRoll === 'function') {
            window.onDiceRoll(result, id);
        }
    }, 600);
}

/**
 * Battle round execution
 * Calculates battle results based on dice rolls
 */
function executeBattleRound() {
    const playerRoll = window.lastPlayerRoll;
    const enemyRoll = window.lastEnemyRoll;
    
    if (!playerRoll || !enemyRoll) {
        addBattleLogEntry('Roll both dice before attacking!', 'warning');
        return;
    }
    
    // In production, this would make an HTMX call to the backend
    // For now, just log the rolls
    addBattleLogEntry(`You rolled ${playerRoll}, Enemy rolled ${enemyRoll}`, 'info');
    
    // Reset rolls after use
    window.lastPlayerRoll = null;
    window.lastEnemyRoll = null;
    
    // In production: 
    // htmx.ajax('POST', '/adventure/battle/attack', {
    //     values: { playerRoll, enemyRoll, pageId: currentPageId },
    //     target: '#battle-section'
    // });
}

/**
 * Add entry to battle log
 * @param {string} message - Log message
 * @param {string} type - Message type (info, warning, success, error)
 */
function addBattleLogEntry(message, type = 'info') {
    const battleLog = document.getElementById('battle-log');
    if (!battleLog) return;
    
    const entry = document.createElement('p');
    entry.className = 'text-sm';
    
    // Style based on type
    switch (type) {
        case 'warning':
            entry.classList.add('text-yellow-400');
            break;
        case 'success':
            entry.classList.add('text-green-400');
            break;
        case 'error':
            entry.classList.add('text-destructive');
            break;
        default:
            entry.classList.add('text-muted-foreground');
    }
    
    entry.textContent = message;
    battleLog.appendChild(entry);
    
    // Auto-scroll to bottom
    battleLog.scrollTop = battleLog.scrollHeight;
}

/**
 * Dashboard tab switching
 */
function switchDashboardTab(tabName) {
    switchTab(tabName);
}

/**
 * Initialize page
 * Run when DOM is ready
 */
document.addEventListener('DOMContentLoaded', function() {
    console.log('Fighting Fantasy - The Forest of Doom initialized');
    
    // Initialize any default states
    window.lastPlayerRoll = null;
    window.lastEnemyRoll = null;
});

/**
 * HTMX event listeners
 * These fire when HTMX loads content
 */
if (typeof htmx !== 'undefined') {
    // After HTMX swaps content
    document.body.addEventListener('htmx:afterSwap', function(event) {
        console.log('Content swapped:', event.detail.target);
    });
    
    // Before HTMX sends request
    document.body.addEventListener('htmx:configRequest', function(event) {
        console.log('Request to:', event.detail.path);
    });
}

// Toast: dismiss button + auto-dismiss
function initToast() {
  const toast = document.getElementById('toast-danger');
  if (!toast) return;

  // Enter animation: start hidden, then reveal
  toast.classList.add('opacity-0', 'translate-y-2');
  requestAnimationFrame(() => {
    requestAnimationFrame(() => { // double rAF forces a paint between add and remove
      toast.classList.remove('opacity-0', 'translate-y-2');
    });
  });

  // Dismiss buttons
  toast.querySelectorAll('[data-dismiss-target]').forEach(btn => {
    btn.addEventListener('click', () => dismissToast(toast));
  });

  // Auto-dismiss
  setTimeout(() => dismissToast(toast), 6000);
}

function dismissToast(el) {
  if (!el) return;
  el.classList.add('opacity-0', 'translate-y-2');
  el.addEventListener('transitionend', () => el.remove(), { once: true });
}

document.addEventListener('htmx:afterSwap', (e) => {
  if (e.detail.target.id === 'auth-error-container') {
    initToast();
  }
});

document.addEventListener('DOMContentLoaded', initToast);
